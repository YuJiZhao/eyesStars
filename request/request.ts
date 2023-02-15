import config from "@/config";

interface RespInterface {
    code: number;
    msg: string;
    [propName: string]: any;
}

const fetch = (url: string, options?: any): Promise<RespInterface> => {
    return new Promise((resolve, reject) => {
        useFetch(url, {
            baseURL: config.apiBaseUrl,
            ...options
        }).then(({ data, error }) => {
            if (error.value) {
                // TODO: 弹窗写好后放弹窗里
                console.log("应该是服务器崩了,或者程序出bug了");
                console.log(error.value);
                return;
            }
            resolve(data.value as RespInterface);
        });
    });
}

export default new class Http {
    get(url: string, params?: any): Promise<RespInterface> {
        return fetch(url, { method: 'get', params });
    }
    post(url: string, params?: any): Promise<RespInterface> {
        return fetch(url, { method: 'post', params });
    }
    put(url: string, body?: any): Promise<RespInterface> {
        return fetch(url, { method: 'put', body });
    }
    delete(url: string, body?: any): Promise<RespInterface> {
        return fetch(url, { method: 'delete', body });
    }
}