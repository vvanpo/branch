import html from "./button.html"
import style from "./button.scss"
import { shadow, template } from "./template.ts"

const tmpl = template(html, style)

export default class Button extends HTMLElement {
    constructor() {
        super()
        shadow(this, tmpl)
    }
}

customElements.define("titian-button", Button)
