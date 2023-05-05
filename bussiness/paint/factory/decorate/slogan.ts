import { slogan } from "@/bussiness/paint/product";
import { LayoutModeEnum } from "@/constant/enum";
import config from "@/config";

// 标语工厂
export default (mode: LayoutModeEnum, length: number, height: number, zr: any, text: string, fontUrl: string, fontFamily: string) => {
    let position = mode == LayoutModeEnum.NORMAL ? normalPosition(length, height) : flipPosition(length, height);
    slogan(
        zr, 
        position.x,
        position.y,
        mode == LayoutModeEnum.NORMAL ? 0 : 90,
        text, 
        fontUrl, 
        fontFamily,
        mode == LayoutModeEnum.NORMAL ? config.sloganSizeNormal : config.sloganSizeFlip
    );
}

// 正常模式位置
function normalPosition(length: number, height: number) {
    return {
        x: 20,
        y: height - config.sloganSizeNormal - 20
    };
}

// 翻转模式位置
function flipPosition(length: number, height: number) {
    return {
        x: config.sloganSizeFlip + 10,
        y: 10
    };
}