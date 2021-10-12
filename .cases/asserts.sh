## Simple asserts in bash

cli_code_non_zero() {
  echo "---------"
  if [ "$?" -eq "0" ]; then 
    echo "PASS"
  else
    echo "FAIL: command should fail" && exit 1
  fi
  echo "----------"
}

cli_code_zero() {
  echo $?
  echo "---------"
  if [ $? -ne 0 ]; then 
    echo "FAIL: Command should succeed" && exit 1
  else
    echo "PASS"
  fi
  echo "---------"
}