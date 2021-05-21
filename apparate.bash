apparate () {
  # in order to get the exit code from the subshell, we must declare the local var first
  local OUT
  local SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"

  OUT=$("$SCRIPT_DIR/apparate" $@)

  if [[ $? -eq 3 ]]; then
  echo "$OUT"
    cd "$OUT"
  else
    echo -e "$OUT"
  fi
}
