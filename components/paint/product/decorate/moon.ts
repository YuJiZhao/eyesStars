import * as zrender from "zrender";
import config from "@/config";

// 月亮
export default (image: string, x: number, y: number) => {
    let moon = new zrender.Image({
        style: {
            x,
            y,
            image,
            width: config.moonSize,
            height: config.moonSize,
            shadowBlur: 5,
            shadowColor: "rgb(255, 255, 255)",
        },
        zlevel: config.staticZLevel
    });
    moon.on("click", () => {
        alert("我是月亮");
    });
    return moon;
}