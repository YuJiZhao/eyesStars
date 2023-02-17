import * as zrender from "zrender";
import { MonitorInterface } from "@/composables/useMonitor";
import { ProcessInterface } from "@/composables/useProcess";
import { meteorStarFactory, messageStarFactory } from "../factory";
import config from "@/config";

// 绘制临时元素
export let drawProvision = (zr: any, monitor: MonitorInterface, process: ProcessInterface) => {
    let group = new zrender.Group();
    drawMessageStar(group, monitor, process);
    drawMeteorStar(group, monitor);
    zr.add(group);
}

// 绘制寄语星星
function drawMessageStar(group: any, monitor: MonitorInterface, process: ProcessInterface) {
    group.add(messageStarFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height,
        () => {
            process.isShowPublish = false;
            process.isShowNotice = false;
            process.isShowStar = true;
        }
    ));
}

// 绘制流星
function drawMeteorStar(group: any, monitor: MonitorInterface) {
    setInterval(() => {
        if (Math.random() < config.meteorStarChance) {
            let meteorStar = meteorStarFactory(
                monitor.layoutMode,
                monitor.length,
                monitor.height,
                () => group.remove(meteorStar)
            );
            group.add(meteorStar);
        }
    }, config.meteorStarAnimateDuration);
}