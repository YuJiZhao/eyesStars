export interface ContextInterface {
    appName: string;
    appNameEN: string;
    appTitle: string;
    appVersion: string;
    siteUrl: string;
    metaDesc: string;
    metaKeyWord: string;
    favicon: string;
    slogan: string;
    sloganFontFamily: string;
    sloganFontUrl: string;
    aesKey: string;
    moonUrl: string;
    notice: string;
    githubAddress: string;
    authorSite: string;
    defaultAvatar: string;
    messageMaxLen: string;
}

export const useContext = () => useState<Partial<ContextInterface>>("context", () => {
    return {};
})