import html from "./button.html"
import style from "./button.scss"
import { shadow, template } from "./template.ts"

const tmpl = template(html, style)
const attrs = ['disabled', 'href']

export default class Button extends HTMLElement {
    constructor() {
        super()
        shadow(this, tmpl)
    }

    connectedCallback() {

    }

    static get observedAttributes() {
        return attrs
    }

    attributeChangedCallback(name, old, val) {
        if (attrs.includes(name)) {
            this[name] = val
        }
    }
}

customElements.define("titian-button", Button)
