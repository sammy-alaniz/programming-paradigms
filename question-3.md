# Question

3. Write an OCaml function that takes as input a list and returns a
list such that each element is computed as an exponent (index+1) of
the original value.  (You should not use any library function in to
solve this problem.)

input: [ 1; 2; 6 ]
output: [1; 4; 216] (* which is: [ 1^1; 2^2; 6^3 ] *)

# Answer
- This function uses a helper function `aux` with an index `i` that starts from `1`. 
- For each element `x` in the input list, it computes `x` raised to the power of `i` and then recursively calls itself with the rest of the list and an incremented index.
- The `float_of_int` function is used to convert the integer index to a float because exponentiation in OCaml requires the exponent to be a float.
```
let exp_list lst =
  let rec aux lst index =
    match lst with
    | [] -> []
    | hd :: tl -> (int_of_float (float_of_int hd ** float_of_int (index + 1))) :: aux tl (index + 1)
  in
  aux lst 0
```
- To test this function with the input [1; 2; 6], you would use it as follows:
```
let () =
  let result = exp_list [1; 2; 6] in
  List.iter (Printf.printf "%d; ") result;
  Printf.printf "\n"
```
- This would give you the output [1.; 4.; 216.], which is [1^1; 2^2; 6^3] in OCaml. 
- Note that the input list elements and the base of the exponent need to be of type float since OCaml's ** operator is defined for floating-point numbers. ​


#### Python equivalnet:
```
def exponentiate_list(input_list):
    """ 
    Takes an input list and returns a list such that each element is computed as an exponent (index+1) 
    of the original value.
    """
    return [input_list[i] ** (i + 1) for i in range(len(input_list))]

# Test the function with the provided input
input_list = [1, 2, 6]
output_list = exponentiate_list(input_list)
output_list
```