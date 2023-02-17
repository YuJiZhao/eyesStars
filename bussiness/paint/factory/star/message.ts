import * as zrender from "zrender";
import { messageStar } from "@/bussiness/paint/product";
import { LayoutModeEnum } from "@/constant/enum";
import config from "@/config";
import { getDoubleRandom } from "@/utils/help";

// 寄语星星工厂
export default (mode: LayoutModeEnum, length: number, height: number, clickFunc: () => void) => {
    let group = new zrender.Group();
    let starNum = config.messageStarNum * length * height;
    for (let i = 0; i < starNum; i++) {
        let position = mode == LayoutModeEnum.NORMAL ? normalPosition(length, height) : flipPosition(length, height);
        let star = messageStar(
            position.x,
            position.y,
            () => {
                clickFunc();
                group.remove(star);
            }
        );
        group.add(star);
    }
    return group;
}

// 正常模式位置
function normalPosition(length: number, height: number) {
    return {
        x: getDoubleRandom(0, length),
        y: getDoubleRandom(0, height),
    };
}

// 翻转模式位置
function flipPosition(length: number, height: number) {
    return {
        x: getDoubleRandom(0, height),
        y: getDoubleRandom(0, length),
    };
}