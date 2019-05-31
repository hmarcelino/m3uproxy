import {Config} from "../model/config";
import * as File from "fs";
import * as JsYaml from "js-yaml";

const YamlConfigLoader = (configFile: string): Config => {
    const yml: string = File.readFileSync(configFile, {encoding: "UTF-8"});
    return JsYaml.safeLoad(yml) as Config;
};

export {YamlConfigLoader};