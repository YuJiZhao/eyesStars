import CryptoJS from "crypto-js";

// 加载字体
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

// 加密
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


// 解密
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

// 获取指定范围内的随机小数
export function getDoubleRandom(min: number, max: number) {
    return min + (max - min) * Math.random();
}

// 获取指定范围内的随机整数
export function getIntRandom(min: number, max: number) {
    return min + Math.round((max - min) * Math.random());
}