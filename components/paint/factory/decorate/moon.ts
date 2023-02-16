import { moon } from "@/components/paint/product";
import { LayoutModeEnum } from "@/constant/enum";
import config from "@/config";

// 月亮工厂
export default (mode: LayoutModeEnum, length: number, height: number, image: string) => {
    let position = mode == LayoutModeEnum.NORMAL ? normalPosition(length, height) : flipPosition(length, height);
    return moon(image, position.x, position.y);
}

// 正常模式位置
function normalPosition(length: number, height: number) {
    return {
        x: length - config.moonSize - 30,
        y: 30
    };
}

// 翻转模式位置
function flipPosition(length: number, height: number) {
    return {
        x: height - config.moonSize - 20,
        y: length - config.moonSize - 20
    };
}