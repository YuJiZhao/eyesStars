import * as zrender from "zrender";
import config from "@/config";

// 背景装饰星星
export default () => {
    return new zrender.Star({
        shape: {
            cx: 40,
            cy: 40,
            n: 6, // 6-8
            r: 3, // 3 - 3.5
            r0: 2, // 2 - 2.5
        },
        style: {
            fill: "rgb(145, 245, 145)", // 145起步
            shadowBlur: 2,
            shadowColor: "#FFF",
            opacity: 0.6, // 0.2 - 0.6
        },
        cursor: "default",
        zlevel: config.staticZLevel
    });
}