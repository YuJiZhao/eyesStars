export interface ProcessInterface {
    isShowTip: boolean;
    tipMsg: string;
    isShowNotice: boolean;
    isShowPublish: boolean;
    isShowStar: boolean;
}

export const useProcess = () => useState<ProcessInterface>("process", () => {
    return {
        isShowTip: false,
        tipMsg: "",
        isShowNotice: false,
        isShowPublish: false,
        isShowStar: false
    }
})