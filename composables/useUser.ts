import config from "@/config";

export interface UserInterface {
    isLoggedin: boolean,
    username: string,
    avatar: string,
    email: string
}

export const useUser = () => useState<UserInterface>("user", () => {
    return {
        isLoggedin: false,
        username: config.userInfoDefault.username,
        avatar: "",
        email: config.userInfoDefault.email
    }
})