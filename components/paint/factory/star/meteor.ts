import { meteorStar } from "@/components/paint/product";
import { LayoutModeEnum } from "@/constant/enum";
import config from "@/config";
import { getDoubleRandom } from "@/utils/help";

// 流星工厂
export default (mode: LayoutModeEnum, length: number, height: number, finishCallBack: () => void) => {
    let position = mode == LayoutModeEnum.NORMAL ? normalPosition(length, height) : flipPosition(length, height);
    return meteorStar(
        position.x,
        position.y,
        mode == LayoutModeEnum.NORMAL ? config.meteorStarRotation : config.meteorStarRotation + 90,
        position.toX,
        position.toY,
        position.sX,
        position.sY,
        finishCallBack
    );
}

// 正常模式位置
function normalPosition(length: number, height: number) {
    let x = length * getDoubleRandom(config.meteorStarStartPosition.min, config.meteorStarStartPosition.max);
    return {
        x,
        y: -config.meteorStarHeight,
        toX: x,
        toY: height * 2,
        sX: 10,
        sY: -10 * Math.sqrt(3)
    };
}

// 翻转模式位置
function flipPosition(length: number, height: number) {
    let y = length * getDoubleRandom(config.meteorStarStartPosition.min, config.meteorStarStartPosition.max);
    return {
        x: height + config.meteorStarHeight,
        y: y,
        toX: height + config.meteorStarHeight,
        toY: y + height * 2,
        sX: 10 * Math.sqrt(3),
        sY: 10
    };
}