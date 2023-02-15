import { LayoutModeEnum } from "@/constant/enum";

export interface MonitorInterface {
    isServer: boolean;
    layoutMode: LayoutModeEnum,
    length: number,
    height: number
}

export const useMonitor = () => useState<MonitorInterface>("monitor", () => {
    return {
        isServer: true,
        layoutMode: LayoutModeEnum.NORMAL,
        length: 0,
        height: 0
    }
})