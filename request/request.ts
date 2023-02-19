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
    beforeFetch(options);
    
    return new Promise((resolve, reject) => {
        useFetch(url, {
            baseURL: initData.baseUrl,
            ...options
        }).then(({ data, error, pending }) => {
            if (error.value) {
                if (!initData.isServer) {
                    showTip("应该是服务器崩了,或者程序出bug了");
                }
                reject(error.value);
                return;
            }
            // TODO: 垃圾东西，useFetch没请求完也tm的执行then
            let timer = setInterval(() => {
                if (!pending.value) {
                    resolve(data.value as RespInterface);
                    clearInterval(timer);
                }
            })
        });
    });
}

// 请求拦截器
const beforeFetch = (options: any) => {
    options.headers = {};

    // 添加鉴权请求头
    if (!initData.isServer && localStorage.getItem(config.storageKey.sToken) && localStorage.getItem(config.storageKey.lToken)) {
        options.headers[config.storageKey.sToken] = localStorage.getItem(config.storageKey.sToken);
        options.headers[config.storageKey.lToken] = localStorage.getItem(config.storageKey.lToken);
    }
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