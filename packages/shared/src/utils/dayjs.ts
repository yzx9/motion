import dayjs from "dayjs"
import isBetween from "dayjs/plugin/isBetween"
import isToday from "dayjs/plugin/isToday"

dayjs.extend(isBetween)
dayjs.extend(isToday)

export { dayjs }
