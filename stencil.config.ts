import { Config } from "@stencil/core"
import { sass } from "@stencil/sass"

export const config: Config = {
    namespace: "titian",
    srcDir: "site",
    plugins: [sass()],
}
