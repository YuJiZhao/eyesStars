import * as zrender from "zrender";
import config from "@/config";
import { MonitorInterface } from "@/composables/useMonitor";
import { ProcessInterface } from "@/composables/useProcess";
import { meteorStarFactory, messageStarFactory } from "../factory";
import { switchStar } from "@/bussiness/process"; 

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
        () => switchStar()
    ));
}

// 绘制流星
function drawMeteorStar(group: any, monitor: MonitorInterface) {
    setInterval(() => {
        // 如果当前页被折叠，浏览器为了节约性能会停止执行动画，当用户切回当前页后，才集中执行
        // 为防止用户长时间折叠页面后再进入出现多个流星的现象，绘制流星前需要判断是否在当前页
        if (!document.hidden && Math.random() < config.meteorStarChance) {
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