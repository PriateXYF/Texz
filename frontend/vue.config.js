module.exports = {
    css: {
        loaderOptions: {
            less: {
                lessOptions: {
                    modifyVars: {
                        // 直接覆盖变量
                        'white': '#000',
                        'text-color': '#111',
                        'cell-background-color': '#eee',
                        'button-default-background-color': '#444',
                        // 或者可以通过 less 文件覆盖（文件路径为绝对路径）
                    },
                },
            },
        },
    },
};