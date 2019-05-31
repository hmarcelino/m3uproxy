export interface Channel {
    url: string
    extraInfo: string

    channelKey(): string
}