import * as zrender from "zrender";
import config from "@/config";

// 月亮
export default (image: string, x: number, y: number, clickFunc: () => void) => {
    let moon = new zrender.Image({
        style: {
            x,
            y,
            image,
            width: 70,
            height: 70,
            shadowBlur: 5,
            shadowColor: "rgb(255, 255, 255)",
        },
        zlevel: config.staticZLevel
    });
    moon.on("click", clickFunc);
    return moon;
}