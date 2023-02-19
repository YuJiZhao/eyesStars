import { http } from "@/request/request";

// 初始化网站配置信息(服务端请求)
export const initSite = async () => {
    return await http.get("/context/initSite");
}

// 初始化用户信息(客户端请求)
export const initUser = async () => {
    return await http.get("/user/userInfoGet");
}

// 批量获取星星(客户端请求)
interface StarsGetInterface {
    ids?: string;
}
export const getStars = async (req: StarsGetInterface) => {
    return await http.get("/star/starsGet", req);
}

// 发布星星(客户端请求)
interface StarAddInterface {
    content: string;
    name?: string;
    anonymous?: boolean;
}
export const addStar = async (req: StarAddInterface) => {
    return await http.post("/star/starAdd", req);
}