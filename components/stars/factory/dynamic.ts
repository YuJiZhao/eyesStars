import * as zrender from "zrender";
import { messageStar, meteorStar } from "@/components/stars/product";
import config from "@/config";

// 绘制动态元素
export let drawDynamic = () => {
    let group = new zrender.Group();
    let zlevel = config.dynamicZLevel;

    group.add(meteorStar(zlevel));
    group.add(messageStar(zlevel));

    return group;
}