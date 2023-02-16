import * as zrender from "zrender";
import { MonitorInterface } from "@/composables/useMonitor";
import { ContextInterface } from "@/composables/useContext";
import { meteorStarFactory, messageStarFactory } from "@/components/paint/factory";

// 绘制临时元素
export let drawProvision = (zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>) => {
    let group = new zrender.Group();
    drawMessageStar(group, monitor);
    drawMeteorStar(group, monitor);
    zr.add(group);
}

// 绘制寄语星星
function drawMessageStar(group: any, monitor: MonitorInterface) {
    group.add(messageStarFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height
    ));
}

// 绘制流星
function drawMeteorStar(group: any, monitor: MonitorInterface) {
    let meteorStar = meteorStarFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height,
        () => group.remove(meteorStar)
    );
    group.add(meteorStar);
}

// 绘制tip

// 绘制寄语弹窗

// 绘制公告弹窗