func addBinary(a string, b string) string {
    // Выбираем самую длинную и короткую строку по их длине
    max_num := a
    min_num := b
    if len(b) > len(a) {
        max_num = b
        min_num = a
    }
    tot := ""
    flag := false
    
    // Считаем разницу в длинах, чтобы правильно сдвинуть индекс короткой строки
    diff := len(max_num) - len(min_num)

    for i := len(max_num) - 1; i >= 0; i-- {
        
        // Вычисляем j (индекс для короткой строки) и проверяем, что он существует
        if j := i - diff; j >= 0 { 
            if flag == true {
                // Заменили min_num[i] на min_num[j]
                if (max_num[i] == '1' && min_num[j] == '0') || (max_num[i] == '0' && min_num[j] == '1') {
                    tot = "0" + tot
                } else if max_num[i] == '1' && min_num[j] == '1'{
                    tot = "1" + tot
                }else {
                    tot = "1" + tot
                    flag = false
                }
            } else {
                // Заменили min_num[i] на min_num[j]
                if (max_num[i] == '1' && min_num[j] == '0') || (max_num[i] == '0' && min_num[j] == '1') {
                    tot = "1" + tot
                } else if max_num[i] == '1' && min_num[j] == '1'{
                    tot = "0" + tot
                    flag = true
                }else  {
                    tot = "0" + tot
                }
            }
        } else { 
            // Этот блок срабатывает, когда короткая строка слева уже закончилась (j < 0)
            if flag == true {
                if max_num[i] == '0' {
                    tot = "1" + tot
                    flag = false
                } else {
                    tot = "0" + tot
                }
            } else {
                tot = string(max_num[i]) + tot
            }
        }
    }
    if flag == true{
        tot = "1" + tot
    }
    return tot 
}