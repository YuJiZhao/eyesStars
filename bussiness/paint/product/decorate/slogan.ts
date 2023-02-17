import * as zrender from "zrender";
import { loadFont } from "@/utils/help";
import config from "@/config";

// 标语
export default (zr: any, x: number, y: number, rotation: number, text: string, fontUrl: string, fontFamily: string, fontSize: number) => {
    loadFont(fontUrl, fontFamily, () => {
      let slogan = new zrender.Text({
        style: {
          x,
          y,
          text,
          fill: "#fff",
          fontSize,
          fontFamily,
        },
        rotation: - rotation / 180 * Math.PI,
        originX: x,
        originY: y,
        cursor: "default",
        zlevel: config.staticZLevel
      });
      zr.add(slogan);
    });
}