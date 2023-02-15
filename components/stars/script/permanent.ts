import * as zrender from "zrender";
import { decorateStar, moon, slogan } from "@/components/stars/product";
import { MonitorInterface } from "@/composables/useMonitor";

// 绘制常驻元素
export let drawPermanent = (monitor: MonitorInterface) => {
    let group = new zrender.Group();
    return group;
}