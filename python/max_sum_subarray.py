

def get_user_input_intList_size()->(list[int], int):
  list_str = input("Enter the list of intergers comma seperated:")
  size = input("Enter the size of subarray:")

  int_list = [int(i) for i in list_str.split(",")]
  return int_list, int(size)


def max_sum_subarray(list, size) -> int :
  if size < 0 or len(list) < size :
    return 0

  window_sum = sum(list[:size])
  max_sum = window_sum

  for i in range(size,len(list)):
    window_sum += list[i] - list[i-size]
    max_sum = max(window_sum, max_sum)

  return max_sum

print(max_sum_subarray([2, 1, 5, 1, 3, 2], 3))

list, size = get_user_input_intList_size()
print(max_sum_subarray(list, size))