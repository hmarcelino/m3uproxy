export interface Config {
    m3uUrlFile: string,
    tempDir: string,
    sourceM3uFile: string,
    proxyM3uFile: string
    nginx: NginxConfig
    logging: Logging
}

export interface NginxConfig {
    baseProxyUrl: string
    wwwRootDir: string
    logDir: string
    templateFile: string
    configFile: string
    reload: string
}

export interface Logging {
    enabled: boolean
    file: string
}