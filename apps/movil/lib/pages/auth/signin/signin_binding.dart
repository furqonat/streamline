

import 'package:get/get.dart';

import 'signin_controller.dart';

class SigninBinding implements Bindings {
  @override
  void dependencies() {
    Get.lazyPut(
      () => SigninController(),
    );
  }
  
}