import * as File from "fs";
import * as Crypto from "crypto";
import * as ChildProcess from "child_process";
// custom
import {M3uFileParser} from "./modules/m3u-file-parser";
import {YamlConfigLoader} from "./modules/config-loader";
// tools
import Axios from "axios"
import * as Mustache from "mustache";
import * as ParseArgs from "minimist";
import {ConsoleLogger, FileLogger, LoggingDispacther} from "./modules/logging";

const args = ParseArgs(process.argv.splice(2), {
    alias: {f: "file"}
});

if (!args.file) {
    console.log(`Missing config file:\nUsage: yarn run run <yaml-config-file-path>`);
    process.exit(1);
}

//********************************
// Initial setup
//********************************

const config = YamlConfigLoader(args.file);

let loggers = [new ConsoleLogger()];
if (config.logging) {
    loggers.push(new FileLogger(config.logging.file))
}

const logging = new LoggingDispacther(loggers);

//********************************
// Proxy process
//********************************

logging.write("=============================================");
logging.write(new Date().toISOString());

Axios({url: config.m3uUrlFile, method: "get"})
    .then(resp => {
        const newM3uFileContent = resp.data;

        if (File.existsSync(`${config.tempDir}/${config.sourceM3uFile}`)) {
            const oldM3uMd5 = Crypto.createHash("md5")
                .update(File.readFileSync(`${config.tempDir}/${config.sourceM3uFile}`))
                .digest("hex");

            const mewM3uMd5 = Crypto.createHash("md5")
                .update(newM3uFileContent)
                .digest("hex");

            if (oldM3uMd5 === mewM3uMd5) {
                throw new Error("M3U file don\"t have any differences");
            }
        }

        File.writeFileSync(`${config.tempDir}/${config.sourceM3uFile}`, newM3uFileContent);
        return new M3uFileParser(newM3uFileContent)
    })
    .then(m3uFileParser => {

        //********************************
        // Build the new m3u file pointing
        // to the local server
        //********************************

        const newM3uFileContent =
            "#EXTM3U\n" +
            m3uFileParser.channels
                .map(channel => `${channel.extraInfo}\n${config.nginx.baseProxyUrl}${channel.channelKey()}`)
                .join("\n");

        File.writeFileSync(`${config.nginx.wwwRootDir}/${config.proxyM3uFile}`, newM3uFileContent);
        return m3uFileParser;
    })
    .then((m3uFileParser) => {

        //********************************
        // Build nginx template file
        //********************************

        const nginxTmpl = File.readFileSync(config.nginx.templateFile, "UTF-8");
        const nginxConfig = Mustache.render(
            nginxTmpl, {
                rootDir: config.nginx.wwwRootDir,
                logDir: config.nginx.logDir,
                channels: m3uFileParser.channels
            });

        File.writeFileSync(`${config.nginx.configFile}`, nginxConfig);

        if (config.nginx.reload) {
            ChildProcess.exec("nginx -s reload")
        }
    })
    .catch((err: any) => {
        logging.write(err.toString());
        process.exit(10);
    })
    .finally(() => logging.write("process finished"));

