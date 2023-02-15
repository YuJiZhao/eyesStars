import * as zrender from "zrender";

// 寄语星星
export default (zlevel: number) => {
    let star = new zrender.Star({
        shape: {
            cx: 80,
            cy: 80,
            n: 8, //
            r: 5, //
            r0: 2, //
        },
        style: {
            fill: "rgb(155, 145, 145)", // 145起步
            shadowBlur: 3,
            shadowColor: "rgb(15, 245, 145)",
            opacity: 0.8, // 0.2 - 0.6
        },
        zlevel
    });
    star.on("click", () => {
        console.log(12);
        star.off("click");
    });
    return star;
}