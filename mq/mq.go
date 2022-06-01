package mq

import (
	"encoding/json"
	"errors"
	"strings"

	"github.com/streadway/amqp"
)

var conn *amqp.Connection
var channel *amqp.Channel
var topics string
var nodes string
var hasMQ bool = false

type Reader interface {
	Read(msg *string) (err error)
}

// 初始化 参数格式：amqp://用户名:密码@地址:端口号/host
func SetupRMQ(rmqAddr string) (err error) {
	if channel == nil {
		conn, err = amqp.Dial(rmqAddr)
		if err != nil {
			return err
		}

		channel, err = conn.Channel()
		if err != nil {
			return err
		}

		hasMQ = true
	}
	return nil
}

// 是否已经初始化
func HasMQ() bool {
	return hasMQ
}

// 测试连接是否正常
func Ping() (err error) {

	if !hasMQ || channel == nil {
		return errors.New("RabbitMQ is not initialize")
	}

	err = channel.ExchangeDeclare("ping.ping", "topic", false, true, false, true, nil)
	if err != nil {
		return err
	}

	msgContent := "ping.ping"

	err = channel.Publish("ping.ping", "ping.ping", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte(msgContent),
	})

	if err != nil {
		return err
	}

	err = channel.ExchangeDelete("ping.ping", false, false)

	return err
}

// 发布消息
func Publish(topic, node string, msg interface{}) (err error) {

	if topics == "" || !strings.Contains(topics, topic) {
		//name :交换器的名称
		//	kind : 交换器的类型，常见的有direct,fanout,topic等，
		//			Direct Pattern （此模式不需要配置交换机）
		//			Fanout Pattern ( 类似于广播一样，将消息发送给和他绑定的队列 )
		//			Topic Pattern ( 绑定交换机时可以做匹配。 #：表示零个或多个单词。*：表示一个单词 )
		//			Header Pattern ( 带有参数的匹配规则 )
		//  durable :设置是否持久化。durable设置为true时表示持久化，反之非持久化.持久化可以将交换器存入磁盘，在服务器重启的时候不会丢失相关信息。
		//	autoDelete：设置是否自动删除。autoDelete设置为true时，则表示自动删除。自动删除的前提是至少有一个队列或者交换器与这个交换器绑定，之后，所有与这个交换器绑定的队列或者交换器都与此解绑。不能错误的理解—当与此交换器连接的客户端都断开连接时，RabbitMq会自动删除本交换器
		//	internal： 设置是否内置的。如果设置为true，则表示是内置的交换器，客户端程序无法直接发送消息到这个交换器中，只能通过交换器路由到交换器这种方式。
		// noWait:是否等待服务器返回
		//arguments:其它一些结构化的参数，比如：alternate-exchange
		err = channel.ExchangeDeclare(topic, "topic", true, false, false, true, nil)
		if err != nil {
			return err
		}
		topics += "  " + topic + "  "
	}

	//routingKey：路由键，#匹配0个或多个单词，*匹配一个单词，在topic exchange做消息转发用
	//mandatory：
	//	true：如果exchange根据自身类型和消息routeKey无法找到一个符合条件的queue，
	//那么会调用basic.return方法将消息返还给生产者。
	//	false：出现上述情形broker会直接将消息扔掉
	//mandatory标志告诉服务器至少将该消息route到一个队列中，否则将消息返还给生产者；
	//immediate标志告诉服务器如果该消息关联的queue上有消费者，则马上将消息投递给它，
	//如果所有queue都没有消费者，直接把消息返还给生产者，不用将消息入队列等待消费者了。

	mbytes, err := json.Marshal(msg)
	err = channel.Publish(topic, node, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        mbytes,
	})

	return err
}

// 监听接收到的消息
func Receive(topic, node string, reader func(msg []byte)) (err error) {
	if topics == "" || !strings.Contains(topics, topic) {
		err = channel.ExchangeDeclare(topic, "topic", true, false, false, true, nil)
		if err != nil {
			return err
		}
		topics += "  " + topic + "  "
	}
	if nodes == "" || !strings.Contains(nodes, node) {
		//参数1(String queue):指定队列的名称，如果队列不存在，则会自动创建一个队列，名称为传入的queue；
		//参数2(boolean druable):定义队列特性是否需要持久化，如果为false，即在下一次启动mq时，队列会被删除。如果为true，下一次启动时，队列仍会存在，但需要注意的是，队列中的消息不会保留，下一次启动时，队列为空。                                                         如果需要消息的持久化，可以在利用basicPublish函数传递消息时，指明MessageProperties.PERSISTENT_TEXT_PLAIN，下次启动时，会恢复队列中的消息。
		//参数3(boolean exclusive):是否独占队列，如果为true，则该队列只能与当前的通道绑定，其他的通道访问不了该队列
		//参数4(boolean autoDelete):是否自动删除，在消费者消费完队列中的数据并与该队列连接断开时，是否要删除该队列
		//参数5(Map<String, Object>arguments)：额外参数设置
		_, err = channel.QueueDeclare(node, true, false, false, true, nil)
		if err != nil {
			return err
		}
		//queue 队列名称
		//exchange 交换器名称
		//routingKey 路由key
		//arguments 其它的一些参数
		err = channel.QueueBind(node, topic, topic, true, nil)
		if err != nil {
			return err
		}
		nodes += "  " + node + "  "
	}
	//订阅消息并消费
	//参数：queue
	//含义：所订阅的队列
	//参数：autoAck
	//含义：是否开启自动应答，默认是开启的，如果需要手动应答应该设置为false
	//注意：为了确保消息一定被消费者处理，rabbitMQ提供了消息确认功能，
	//就是在消费者处理完任务之后，就给服务器一个回馈，服务器就会将该消息删除，
	//如果消费者超时不回馈，那么服务器将就将该消息重新发送给其他消费者，
	//当autoAck设置为true时，只要消息被消费者处理，不管成功与否，服务器都会删除该消息，
	//而当autoAck设置为false时，只有消息被处理，且反馈结果后才会删除。
	//参数：callback
	//含义：接收到消息之后执行的回调方法
	//参数：exclusive: 是否独占。
	//noLocal	未使用的参数。
	//noWait	如果为 True，那么其他连接尝试修改该队列，将会触发异常。
	msgs, err := channel.Consume(node, "", true, false, false, false, nil)
	if err != nil {
		return err
	}

	go func() {
		//fmt.Println(*msgs)
		for d := range msgs {
			reader(d.Body)
		}
	}()

	return err
}

// 关闭连接
func Close() {
	channel.Close()
	conn.Close()
	hasMQ = false
}
