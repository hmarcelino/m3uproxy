import * as File from "fs";

export interface Logger {
    write(text: string): void
}

export class ConsoleLogger implements Logger {
    write(text: string): void {
        console.log(text);
    }
}

export class FileLogger implements Logger {
    private readonly _logFile: string;

    constructor(logFile: string) {
        if (!File.existsSync(logFile)) {
            File.writeFileSync(logFile, "");

        }
        this._logFile = logFile;
    }

    write(text: string): void {
        if (!text) return;

        File.appendFileSync(this._logFile, text);

        if (!text.endsWith("\n")) {
            File.appendFileSync(this._logFile, "\n");
        }
    }
}

export class LoggingDispacther {

    private readonly _loggers: Logger[] = [];

    constructor(loggers: Logger[]) {
        if (loggers !== null) {
            this._loggers = loggers;
        }
    }

    write(text: string): void {
        for (const loggerElement of this._loggers) {
            loggerElement.write(text);
        }
    }
}