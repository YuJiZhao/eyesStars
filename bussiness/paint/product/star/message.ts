import * as zrender from "zrender";
import config from "@/config";
import { getDoubleRandom, getIntRandom } from "@/utils/help";

// 寄语星星
export default (x: number, y: number, clickFunc: () => void) => {
    let initOpacity = getDoubleRandom(config.messageStarOpacity.min, config.messageStarOpacity.max);
    let star = new zrender.Star({
        shape: {
            cx: x,
            cy: y,
            n: getIntRandom(config.messageStarN.min, config.messageStarN.max),
            r: getDoubleRandom(config.messageStarR.min, config.messageStarR.max),
            r0: getDoubleRandom(config.messageStarR0.min, config.messageStarR0.max),
        },
        style: {
            fill: `rgb(
                ${getDoubleRandom(config.decorateStarFill.min, config.decorateStarFill.max)}, 
                ${getDoubleRandom(config.decorateStarFill.min, config.decorateStarFill.max)}, 
                ${getDoubleRandom(config.decorateStarFill.min, config.decorateStarFill.max)}
            )`,
            opacity: initOpacity,
        },
        rotation: getIntRandom(0, 90) / 90 * Math.PI,
        originX: x,
        originY: y,
        zlevel: config.dynamicZLevel
    });
    
    star.animate("style", true)
    .delay(getIntRandom(0, config.messageStarAnimateDuration))
    .when(config.messageStarAnimateDuration, {
        opacity: 0.1
    })
    .when(config.messageStarAnimateDuration * 2, {
        opacity: initOpacity
    })
    .start("exponentialInOut");

    star.on("click", clickFunc);
    return star;
}