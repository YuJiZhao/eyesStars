import * as zrender from "zrender";
import { messageStar, meteorStar } from "@/components/stars/product";
import { MonitorInterface } from "@/composables/useMonitor";

// 绘制临时元素
export let drawProvision = (monitor: MonitorInterface) => {
    let group = new zrender.Group();
    return group;
}