let () = print_endline "Hello, World!"

let exp_list lst = 
  (* line below defines a new function that can be recursivley called "rec", function is named "aux", and expects a list, variable name "lst", and an int of the index, variable name "index" *)
  let rec aux lst index =
    match lst with
    | [] -> []
    | hd :: tl -> (int_of_float (float_of_int hd ** float_of_int (index + 1))) :: aux tl (index + 1)
  in
  aux lst 0

  let () =
  let result = exp_list [1; 2; 6] in
  List.iter (Printf.printf "%d; ") result;
  Printf.printf "\n"
