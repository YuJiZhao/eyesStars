import { http } from "@/request/request";

// 初始化网站配置信息(服务端请求)
export const initSite = async () => {
    return await http.get("/context/initSite");
}

// 进站埋点(客户端请求)
interface VisitTrackInterface {
    package: string;
}
export const visitTrack = async (req: VisitTrackInterface) => {
    await http.get("/track/visit", req);
};

// 登录埋点
interface LoginTrackInterface {
    package: string;
}
export const loginTrack = async (req: LoginTrackInterface) => {
    await http.get("/track/login", req);
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