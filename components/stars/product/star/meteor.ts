import * as zrender from "zrender";

// 流星
export default (zlevel: number) => {
    let meteor = new zrender.Droplet({
        shape: {
            cx: 290,
            cy: 290,
            width: 10,
            height: 25,
        },
        style: {
            fill: "rgb(255, 255, 255)",
            shadowBlur: 5,
            shadowColor: "rgb(253, 219, 215)",
            shadowOffsetX: 0,
            shadowOffsetY: -2
        },
        zlevel
    });
    meteor.on("click", () => {
    });
    return meteor;
}