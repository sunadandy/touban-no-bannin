import { mount, shallowMount  } from '@vue/test-utils'
import CalcMemberInfo from '@/components/module/CalcMemberInfo.vue'
import { format, startOfToday } from 'date-fns'

describe('CalcMemberInfo.vue', () => {
  const today = format(startOfToday(), 'yyyy-MM-dd')
  const wrapper = mount(CalcMemberInfo)

  it('Test _GetNextMonthDay', () => {
    // 第1週の月曜日
    let ret = wrapper.vm._GetNextMonthDay("2024-02-17", 1, 1)
    expect(ret).toBe("2024-03-04")
    // 第1週の金曜日
    ret = wrapper.vm._GetNextMonthDay("2024-02-17", 1, 5)
    expect(ret).toBe("2024-03-01")
    // 第5週の月曜日
    ret = wrapper.vm._GetNextMonthDay("2024-02-17", 5, 1)
    expect(ret).toBe("2024-03-25")
    // 第5週の金曜日
    ret = wrapper.vm._GetNextMonthDay("2024-02-17", 5, 5)
    expect(ret).toBe("2024-03-29")
    // 第3週の水曜
    ret = wrapper.vm._GetNextMonthDay("2024-02-17", 3, 3)
    expect(ret).toBe("2024-03-13")
  })

  it('Test _GetNextMonthSpecifiedDate', () => {
    // 翌月の7日
    let ret = wrapper.vm._GetNextMonthSpecifiedDate(today, 7)
    expect(ret).toBe("2024-03-07")
  })

  it('Test _GetNextDateByIncrementWeek', () => {
    // 翌週
    let ret = wrapper.vm._GetNextDateByIncrementWeek("2024-02-17", 7)
    expect(ret).toBe("2024-02-24")
    // 2週後
    ret = wrapper.vm._GetNextDateByIncrementWeek("2024-02-17", 14)
    expect(ret).toBe("2024-03-02")
    // 3週後
    ret = wrapper.vm._GetNextDateByIncrementWeek("2024-02-17", 21)
    expect(ret).toBe("2024-03-09")
  })

  it('Test _GetNextWorkingDate', () => {
    // 翌日
    let ret = wrapper.vm._GetNextWorkingDate("2024-02-22")
    expect(ret).toBe("2024-02-23")
    // 翌日（金曜日の場合）
    ret = wrapper.vm._GetNextWorkingDate("2024-02-23")
    expect(ret).toBe("2024-02-26")
  })
})