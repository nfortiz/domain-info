const variants = new Map()

variants.set("A+", "success")
variants.set("A", "success")
variants.set("A-", "success")
variants.set("B", "warning")
variants.set("C", "warning")
variants.set("D", "warning")
variants.set("F", "danger")
variants.set("M", "danger")
variants.set("T","danger")

function getGrade(grade) {

    return variants.get(grade) || 'danger'
}

export default getGrade