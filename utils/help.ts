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

// 防抖
export function debounce(fn: () => void, delay = 1000) {
    let timer: NodeJS.Timeout;
    return () => {
        if (timer) {
            clearTimeout(timer);
        }
        timer = setTimeout(() => {
            fn();
        }, delay)
    }
}

// 节流
export function throttle(fn: () => void, delay = 1000) {
    let timer: NodeJS.Timeout | null;
    return () => {
        if(timer) {
            return;
        }
        timer = setTimeout(() => {
            fn();
            timer = null;
        }, delay)
    }
}

// 获取用户浏览器类型
export function browser() {
    let u = navigator.userAgent;
    let bws = [{
        name: 'sgssapp',
        it: /sogousearch/i.test(u)
    }, {
        name: 'wechat',
        it: /MicroMessenger/i.test(u)
    }, {
        name: 'weibo',
        it: !!u.match(/Weibo/i)
    }, {
        name: 'uc',
        it: !!u.match(/UCBrowser/i) || u.indexOf(' UBrowser') > -1
    }, {
        name: 'sogou',
        it: u.indexOf('MetaSr') > -1 || u.indexOf('Sogou') > -1
    }, {
        name: 'xiaomi',
        it: u.indexOf('MiuiBrowser') > -1
    }, {
        name: 'baidu',
        it: u.indexOf('Baidu') > -1 || u.indexOf('BIDUBrowser') > -1
    }, {
        name: '360',
        it: u.indexOf('360EE') > -1 || u.indexOf('360SE') > -1
    }, {
        name: '2345',
        it: u.indexOf('2345Explorer') > -1
    }, {
        name: 'edge',
        it: u.indexOf('Edg') > -1
    }, {
        name: 'ie11',
        it: u.indexOf('Trident') > -1 && u.indexOf('rv:11.0') > -1
    }, {
        name: 'ie',
        it: u.indexOf('compatible') > -1 && u.indexOf('MSIE') > -1
    }, {
        name: 'firefox',
        it: u.indexOf('Firefox') > -1
    }, {
        name: 'safari',
        it: u.indexOf('Safari') > -1 && u.indexOf('Chrome') === -1
    }, {
        name: 'QQbrowser',
        it: u.indexOf('MQQBrowser') > -1 && u.indexOf(' QQ') === -1
    }, {
        name: 'QQ',
        it: u.indexOf('QQ') > -1
    }, {
        name: 'chrome',
        it: u.indexOf('Chrome') > -1 || u.indexOf('CriOS') > -1
    }, {
        name: 'opera',
        it: u.indexOf('Opera') > -1 || u.indexOf('OPR') > -1
    }]
    for (var i = 0; i < bws.length; i++) {
        if (bws[i].it) {
            return bws[i].name;
        }
    }
    return 'other';
};

// 获取用户操作系统类型
export function OS() {
    let u = navigator.userAgent;
    if (!!u.match(/compatible/i) || u.match(/Windows/i)) {
        if (u.indexOf("Windows NT 5.0") > -1 || u.indexOf("Windows 2000") > -1) {
            return "Win2000";
        }
        if (u.indexOf("Windows NT 5.1") > -1 || u.indexOf("Windows XP") > -1) {
            return "WinXP";
        }
        if (u.indexOf("Windows NT 5.2") > -1 || u.indexOf("Windows 2003") > -1) {
            return "Win2003";
        }
        if (u.indexOf("Windows NT 6.0") > -1 || u.indexOf("Windows Vista") > -1) {
            return "WinVista";
        }
        if (u.indexOf("Windows NT 6.1") > -1 || u.indexOf("Windows 7") > -1) {
            return "Win7";
        }
        if (u.indexOf("Windows NT 10") > -1 || u.indexOf("Windows 10") > -1) {
            return "Win10";
        }
    } else if (!!u.match(/Macintosh/i) || u.match(/MacIntel/i)) {
        return "MacOS";
    } else if (!!u.match(/iphone/i) || u.match(/Ipad/i)) {
        return "iOS";
    } else if (u.match(/android/i)) {
        return "Android";
    } else if (u.match(/Ubuntu/i)) {
        return "Ubuntu";
    } else {
        return "other";
    }
};