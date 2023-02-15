import * as zrender from "zrender";

// 月亮
export default (img: string) => {
    let moon = new zrender.Image({
        style: {
            x: 150,
            y: 150,
            image: img,
            width: 70,
            height: 70,
            shadowBlur: 5,
            shadowColor: "rgb(255, 255, 255)",
            opacity: 1,
        },
    });
    moon.on("click", () => {

    });
    return moon;
}