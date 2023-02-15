import CryptoJS from "crypto-js";

export function loadFont(url: string, family: string, callback?: Function) {
    const font = new FontFace(
        family,
        "url(" + url + ")"
    );
    font.load().then(() => {
        document.fonts.add(font);
        if (callback) callback();
    });
};

export function aesEncrypt(str: string, key: string) {
    return CryptoJS.AES.encrypt(
        str,
        CryptoJS.enc.Utf8.parse(key),
        {
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
            iv: CryptoJS.enc.Utf8.parse(key)
        }
    ).toString();
};

export function aesDecrypt(str: string, key: string) {
    return CryptoJS.AES.decrypt(
        str,
        CryptoJS.enc.Utf8.parse(key),
        {
            mode: CryptoJS.mode.CBC,
            padding: CryptoJS.pad.Pkcs7,
            iv: CryptoJS.enc.Utf8.parse(key)
        }
    ).toString(CryptoJS.enc.Utf8);
};