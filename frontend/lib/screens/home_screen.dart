import 'package:flutter/material.dart';

class HomeScreen extends StatelessWidget {
  final key = GlobalKey<FormState>();
  final thbtAmountController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('To The Moon'),
        backgroundColor: Color(0xFF00A6E6),
        // เป็นเส้นแบ่งระหว่าง appbar กับ body
        elevation: 0,
      ),
      body: Stack(
        children: [
          Align(
            alignment: Alignment.topCenter,
            child: Container(
              alignment: Alignment.topCenter,
              // ขนาดหน้าจอ
              height: MediaQuery.of(context).size.height * 0.3,
              width: double.infinity,
              color: Color(0xFF0000CC),
              child: Icon(
                Icons.nights_stay,
                size: 170,
                color: Colors.white,
              ),
            ),
          ),
          Align(
            alignment: Alignment.bottomCenter,
            child: Container(
              height: MediaQuery.of(context).size.height * 0.7,
              width: double.infinity,
              decoration: BoxDecoration(
                color: Color(0xFF99FFCC),
                borderRadius: BorderRadius.only(
                  topLeft: Radius.circular(50),
                  topRight: Radius.circular(50),
                ),
              ),
              child: buildBody(),
            ),
          )
        ],
      ),
    );
  }

  Widget buildBody() => Form(
        key: key,
        child: Container(
          padding: EdgeInsets.all(20),
          // color: Colors.yellow,
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            mainAxisSize: MainAxisSize.max,
            children: [
              buildTHBTInput(),
            ],
          ),
        ),
      );

  Widget buildTHBTInput() => Container(
      margin: EdgeInsets.only(
        left: 100,
        top: 30,
        right: 100,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text('Amount to buy (THBT)'),
          TextFormField(
            controller: thbtAmountController,
            keyboardType: TextInputType.number,
            decoration: InputDecoration(
              prefixIcon: Icon(Icons.attach_money),
              labelText: '0',
              border: OutlineInputBorder(
                borderRadius: BorderRadius.circular(25),
                borderSide: BorderSide(color: Colors.grey),
              ),
            ),
            validator: (value) {
              value ??= '';
              if (value.isEmpty) return 'กรุณากรอกอีเมล';
              // if (!RegExp(r'^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$').hasMatch(value)) return 'อีเมลไม่ถูกต้อง';
            },
          ),
        ],
      ));
}
