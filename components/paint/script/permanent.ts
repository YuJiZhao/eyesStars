import { MonitorInterface } from "@/composables/useMonitor";
import { ContextInterface } from "@/composables/useContext";
import { sloganFactory, moonFactory, decorateStarFactory } from "@/components/paint/factory";

// 绘制常驻元素
export let drawPermanent = (zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>) => {
    drawSlogan(zr, monitor, context);
    drawMoon(zr, monitor, context);
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
function drawMoon(zr: any, monitor: MonitorInterface, context: Partial<ContextInterface>) {
    zr.add(moonFactory(
        monitor.layoutMode,
        monitor.length,
        monitor.height,
        context.moonUrl!
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