import { Component, Prop } from "@stencil/core"

@Component({
    tag: "titian-button",
    styleUrl: "button.scss",
})
export class Button {

    @Prop() href: string

    render() {
        return (
            <button type="button">
                <slot></slot>
            </button>
        )
    }

    private tag() {
        return this.href ? "a" : "button"
    }
}
