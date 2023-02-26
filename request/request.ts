import config from "@/config";
import { showTip } from "@/bussiness/process";

interface RespInterface {
    code: number;
    msg: string;
    [propName: string]: any;
}

let initData = {
    isServer: false,
    baseUrl: ""
}

// 请求封装
const fetch = (url: string, options?: any): Promise<RespInterface> => {
    return new Promise((resolve, reject) => {
        useFetch(url, {
            baseURL: initData.baseUrl,
            ...options,
            onRequest({ options }) {
                options.headers = {};
                if (!initData.isServer && localStorage.getItem(config.storageKey.sToken) && localStorage.getItem(config.storageKey.lToken)) {
                    options.headers[config.storageKey.sToken] = localStorage.getItem(config.storageKey.sToken)!;
                    options.headers[config.storageKey.lToken] = localStorage.getItem(config.storageKey.lToken)!;
                }
            },
            onResponse({ response }) {
                let sToken = response.headers.get(config.storageKey.sToken);
                let lToken = response.headers.get(config.storageKey.lToken);
                if (sToken && lToken) {
                    localStorage.setItem(config.storageKey.sToken, sToken);
                    localStorage.setItem(config.storageKey.lToken, lToken);
                }
                resolve(response._data as RespInterface);
            },
            onResponseError({ response }) {
                showTip("应该是服务器崩了,或者程序出bug了");
                reject(response);
            },
            onRequestError({ response }) {
                showTip("我也不知道出什么问题了");
                reject(response);
            }
        });
    });
}

// 初始化请求配置
export function initHttp(isServer: boolean, baseUrl: string) {
    initData.isServer = isServer;
    initData.baseUrl = baseUrl;
}

// 整合请求方法
export let http = new class Http {
    get(url: string, params?: any): Promise<RespInterface> {
        return fetch(url, { method: 'get', params });
    }
    post(url: string, body?: any): Promise<RespInterface> {
        return fetch(url, { method: 'post', body });
    }
    put(url: string, body?: any): Promise<RespInterface> {
        return fetch(url, { method: 'put', body });
    }
    delete(url: string, body?: any): Promise<RespInterface> {
        return fetch(url, { method: 'delete', body });
    }
}