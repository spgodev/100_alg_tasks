package recursion

/*
Есть дерево сотрудников: у каждого руководителя есть подчинённые, у которых могут быть свои подчинённые.
Нужно найти заданного сотрудника по имени и вывести цепочку подчинения от корня (CEO) до этого сотрудника. Если сотрудника нет — вернуть nil.
*/
type Employee struct {
	Name    string
	Reports []*Employee
}

/*
Анна
/      \
Борис      Виктор
/     \         \
Галина  Дмитрий    Елена
/
Жанна

findPath(tree, "Дмитрий") -> ["Анна","Борис","Дмитрий"]
findPath(tree, "Пётр")    -> nil
*/
func findPath(root *Employee, employeeNameToSearch string) []string {
	if root == nil {
		return nil
	}
	if root.Name == employeeNameToSearch {
		return []string{root.Name}
	}
	for _, report := range root.Reports {
		path := findPath(report, employeeNameToSearch)
		if path != nil {
			return append([]string{root.Name}, path...)
		}
	}
	return nil
}
