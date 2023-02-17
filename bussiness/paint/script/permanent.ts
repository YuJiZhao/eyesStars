import { MonitorInterface } from "@/composables/useMonitor";
import { ContextInterface } from "@/composables/useContext";
import { ProcessInterface } from "@/composables/useProcess";
import { sloganFactory, moonFactory, decorateStarFactory } from "../factory";

// 绘制常驻元素
export let drawPermanent = (zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>, process: ProcessInterface) => {
    drawSlogan(zr, monitor, context);
    drawMoon(zr, monitor, context, process);
    drawDecorateStar(zr, monitor);
}

// 标语
function drawSlogan(zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>) {
    sloganFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height,
        zr,
        context.slogan!,
        context.sloganFontUrl!,
        context.sloganFontFamily!
    );
}

// 月亮
function drawMoon(zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>, process: ProcessInterface) {
    zr.add(moonFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height,
        context.moonUrl!,
        () => {
            process.isShowPublish = false;
            process.isShowStar = false;
            process.isShowNotice = !process.isShowNotice;
        }
    ));
}

// 背景星星
function drawDecorateStar(zr: any, monitor: MonitorInterface) {
    zr.add(decorateStarFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height
    ));
}