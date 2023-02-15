import * as zrender from "zrender";
import { loadFont } from "@/utils/help";

// 标语
export default (zr: any, text: string, fontUrl: string, fontFamily: string) => {
    loadFont(fontUrl, fontFamily, () => {
        var slogan = new zrender.Text({
          style: {
            x: 30,
            y: 400,
            text: text,
            fill: "#fff",
            fontSize: 20,
            fontFamily: fontFamily
          },
          cursor: "default",
        });
        zr.add(slogan);
      });
}