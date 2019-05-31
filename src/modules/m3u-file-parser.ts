import {Channel} from "../model/m3u";

class M3uChannel implements Channel {
    private readonly _url: string;
    private readonly _extraInfo: string;
    private key: string;

    constructor(url: string, extraInfo: string) {
        this._url = url;
        this._extraInfo = extraInfo;
    }

    get url(): string {
        return this._url;
    }

    get extraInfo(): string {
        return this._extraInfo;
    }

    channelKey(): string {
        if (!this.key) {
            this.key = this.url.substring(this.url.lastIndexOf("/") + 1, this.url.length);
        }
        return this.key;
    }
}

export class M3uFileParser {
    private readonly _channelsRecord: Channel[];
    private readonly text: string;

    constructor(text: string) {
        this.text = text;

        this._channelsRecord = M3uFileParser.parseFile(this.text)
    }

    get channels(): Channel[] {
        return Object.values(this._channelsRecord);
    }

    private static parseFile(text: string): Channel[] {
        const channels: M3uChannel[] = [];
        const lines = text.split(/\r?\n/);

        for (let idx = 0; idx < lines.length; idx++) {
            if (!lines[idx].startsWith("#EXTINF")) continue;

            channels.push(new M3uChannel(lines[idx + 1], lines[idx]));
            idx++;
        }

        return channels;
    }
}
