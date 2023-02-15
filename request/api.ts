import Http from "@/request/request";

export const initSite = async () => {
    return await Http.get("/context/initSite");
}
