

import 'package:get/get.dart';
import 'package:movil/routes/route.dart';

class SplashController extends GetxController {
  
  @override
  void onInit() {
    super.onInit();
    Future.delayed(const Duration(seconds: 2)).then((v) {
      // todo Move page
      Get.offAllNamed(PageNames.signIn);
    });
  }
}