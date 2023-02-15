import * as zrender from "zrender";
import { loadFont } from "@/utils/help";
import config from "@/config";

// 标语
export default (zr: any, x: number, y: number, text: string, fontUrl: string, fontFamily: string, fontSize: number) => {
    loadFont(fontUrl, fontFamily, () => {
        zr.add(new zrender.Text({
          style: {
            x,
            y,
            text,
            fill: "#fff",
            fontSize,
            fontFamily
          },
          cursor: "default",
          zlevel: config.staticZLevel
        }));
    });
}