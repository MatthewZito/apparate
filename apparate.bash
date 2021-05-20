tp () {
  local OUT=$(apparate $@)
  
  if [[ $? -eq 3 ]];
    cd "$OUT"
  else 
    echo -e "$OUT"
  fi
}

tp_dev () {
  local OUT=$(go run main.go $@)
  
  if [[ $? -eq 3 ]];
    cd "$OUT"
  else 
    echo -e "$OUT"
  fi
}

tp_dev "$@"