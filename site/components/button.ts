import html from "./button.html"
import style from "./button.scss"

export default class extends HTMLElement {
    constructor() {
        super()
        const shadow = this.attachShadow({mode: "open"})
        shadow.innerHTML = "<style>" + style + "</style>" + html
    }
}
